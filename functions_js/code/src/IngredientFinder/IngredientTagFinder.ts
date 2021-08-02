import { Parser } from 'htmlparser2';
import { ClassNameFinder } from './ClassNameFinder';
import { HtmlSection } from './HtmlSection';
import { IngredientWrapperLocator } from './IngredientWrapperLocator';
import { IngredientLocator } from './IngredientLocator';
import { IngredientSection } from './Models';

export class IngredientTagFinder {
    constructor() {
    }

    static async find(html: string): Promise<IngredientSection[]> {
        const classNames = await ClassNameFinder.getAllClassNames(html)

        // try find a class name that is for each ingredient

        const contents = await IngredientTagFinder.readContents(html)
        const ings = await IngredientLocator.find(contents)

        if (ings.length > 0) {
            return ings
        }

        // otherwise try find a class name that could be the wrapper
        const ingredientWrapperTags = IngredientTagFinder.findIngredientWrapperTags(classNames)
        if (ingredientWrapperTags.length > 0) {
            // - look for list of regular tags
            console.log('looking for wrapper tag', ingredientWrapperTags[0])
            const contents = await IngredientTagFinder.readInsideElement(html, (name, classList) => classList.indexOf(ingredientWrapperTags[0]) > -1)

            const r = IngredientWrapperLocator.find(contents)
            if (r.length === 0) {
                return []
            }
            if (r.length === 1) {
                return r
            }

            console.log('what now')
            console.log(JSON.stringify(r, null, 2))

            return []
        }

        return []
    }

    private static async readContents(html: string): Promise<HtmlSection> {
        try {
            return await IngredientTagFinder.readInsideElement(html, (name, classList) => name === 'body')
        } catch {
            // try reading whole page
            return await IngredientTagFinder.readInsideElement(html, (name, classList) => true)
        }
    }

    private static findIngredientWrapperTags(allTags: { [key: string]: number }): string[] {
        // assumes there is more than one ingredient

        const options: string[] = []

        Object.keys(allTags).forEach(tag => {
            if (tag.indexOf('ingredient') > 0) {
                if (allTags[tag] === 1) {
                    options.push(tag)
                }
            }
        })

        return options
    }

    private static readInsideElement(pageHtml: string, isTopElement: (name: string, classList: string[]) => boolean): Promise<HtmlSection> {
        return new Promise((resolve, reject) => {
            let insideCount = -1
            let htmlSection: HtmlSection | undefined;;
            let currentHtmlSection: HtmlSection | undefined;
            function pushSection(section: HtmlSection) {
                if (typeof htmlSection === 'undefined') {
                    htmlSection = section
                }

                if (typeof currentHtmlSection !== 'undefined') {
                    currentHtmlSection.children.push(section)
                    section.parent = currentHtmlSection
                }
                currentHtmlSection = section
            }
            function closeSection() {
                if (typeof currentHtmlSection !== 'undefined') {
                    currentHtmlSection = currentHtmlSection.parent
                }
            }
            function clean(section: HtmlSection) {
                section.parent = undefined
                section.children.forEach(s => clean(s))
            }

            var parser = new Parser({
                onopentag: (name, attribs) => {
                    const classList = !attribs.class ? [] : attribs.class.split(' ')
                    if (insideCount >= 0) {
                        insideCount++
                        pushSection({ classList, tag: name, children: [] })
                    }
                    if (isTopElement(name, classList)) {
                        insideCount = 0;
                        pushSection({ classList, tag: name, children: [] })
                    }
                },
                ontext: function (text) {
                    if (typeof currentHtmlSection !== 'undefined') {
                        text = text.replace(/\s+/g, ' ')
                        currentHtmlSection.text = text.trim()
                    }
                },
                onclosetag: () => {
                    insideCount--;
                    closeSection()
                },
                onend: () => {
                    if (htmlSection) {
                        clean(htmlSection)
                        resolve(htmlSection);
                    }
                    reject('not found')
                },
                onerror: err => {
                    reject(err)
                }
            }, { decodeEntities: true });

            parser.write(pageHtml);
            parser.end();
        })
    }
}
