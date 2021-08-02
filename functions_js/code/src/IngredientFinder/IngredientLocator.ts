import { Parser } from 'htmlparser2';
import { IngredientSection } from './Models';
import { HtmlSection } from './HtmlSection';

type FoundIngredientSection = IngredientSection & { foundInClass: string, foundParentTag?: string, foundParentClassList?: string[] }

export class IngredientLocator {

    static async find(contents: HtmlSection): Promise<IngredientSection[]> {
        const { result, sections } = IngredientLocator.doesSectionContainIngredientList(contents);

        if (result) {
            let foundParentTag = sections[0].foundParentTag
            let foundParentClassList = (sections[0].foundParentClassList || []).sort().join('|')
            let foundInClass = sections[0].foundInClass
            let allsame = true

            sections.forEach(s => {
                if (!allsame) return
                if (s.foundInClass !== foundInClass) {
                    allsame = false
                }
                if (s.foundParentTag !== foundParentTag) {
                    allsame = false
                }
                if ((s.foundParentClassList || []).sort().join('|') !== foundParentClassList) {
                    allsame = false
                }
            })

            if (allsame) {
                return IngredientLocator.findAllOfType(contents, foundParentTag, foundParentClassList, foundInClass)
            }

            return sections
        }

        return []
    }

    static doesSectionContainIngredientList(section: HtmlSection): { result: boolean, sections: FoundIngredientSection[] } {

        if (section.children.length === 0) {
            return { result: false, sections: [] }
        }

        const childrenResults = section.children.map(c => IngredientLocator.doesSectionContainIngredientList(c))
        const childrenFound = childrenResults.filter(r => r.result)

        if (childrenFound.length > 0) {
            const sections = childrenFound.reduce((acc, curr) => {
                const mapped = curr.sections.map(s => {
                    if (!s.foundParentTag) {
                        s.foundParentTag = section.tag
                        s.foundParentClassList = section.classList
                    }
                    return s
                })
                acc.push(...mapped)
                return acc
            }, [] as FoundIngredientSection[])
            return { result: true, sections }
        }

        let classNames = {}
        function pushClassName(className: string) {
            if (/ingredient/.exec(className) === null) return

            if (typeof classNames[className] === 'undefined') {
                classNames[className] = 0
            }
            classNames[className]++
        }
        section.children.forEach(c => {
            if (!c.classList) return

            c.classList.forEach(className => pushClassName(className))
        })

        const classNameKeys = Object.keys(classNames)
        if (classNameKeys.length > 0) {
            let highestClassCount = -1
            let highestClass = ''
            classNameKeys.forEach(cn => {
                if (classNames[cn] > highestClassCount) {
                    highestClassCount = classNames[cn]
                    highestClass = cn
                }
            })

            if (highestClassCount > 1) {
                // do something here?

                let ingredients: string[] = []
                section.children.forEach(c => {
                    if (c.classList.indexOf(highestClass) > -1) {
                        const text = IngredientLocator.readText(c)
                        ingredients.push(text)
                    }
                })

                return { result: true, sections: [{ ingredients, foundInClass: highestClass }] }
            }
        }

        return { result: false, sections: [] }
    }

    private static findAllOfType(section: HtmlSection, parentTag: string, parentClassList: string, className: string): IngredientSection[] {
        const sections: IngredientSection[] = []

        section.children.forEach(c => {
            const foundInChild = IngredientLocator.findAllOfType(c, parentTag, parentClassList, className)
            sections.push(...foundInChild)
        })

        if (section.tag === parentTag && (section.classList || []).sort().join('|') === parentClassList) {
            let ingredients: string[] = []

            function findAllInSection(subsection: HtmlSection) {
                subsection.children.forEach(c => {
                    if (c.classList.indexOf(className) > -1) {
                        const text = IngredientLocator.readText(c)
                        ingredients.push(text)
                    }

                    findAllInSection(c)
                })
            }

            findAllInSection(section)

            sections.push({ ingredients })
        }

        return sections
    }

    private static readText(section: HtmlSection): string {
        let text = (section.text || '').trim() + (section.children || []).map(c => IngredientLocator.readText(c)).join(' ')
        text = text.replace(/\s,/g, ',')
        return text.trim()
    }
}
