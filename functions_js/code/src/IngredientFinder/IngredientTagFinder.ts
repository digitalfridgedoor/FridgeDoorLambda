import { Parser } from 'htmlparser2';
import { ClassNameFinder } from './ClassNameFinder';
import { HtmlSection } from './HtmlSection';
import { IngredientWrapperLocator } from './IngredientWrapperLocator';
import { IngredientLocator } from './IngredientLocator';
import { IngredientSection } from './Models';
import { HtmlReader } from './HtmlReader';

export class IngredientTagFinder {
    constructor() {
    }

    static async find(html: string): Promise<IngredientSection[]> {
        const classNames = await ClassNameFinder.getAllClassNames(html)

        // try find a class name that is for each ingredient

        const contents = await HtmlReader.readContents(html)
        const ings = await IngredientLocator.find(contents)

        if (ings.length > 0) {
            return ings
        }

        // otherwise try find a class name that could be the wrapper
        const ingredientWrapperTags = IngredientTagFinder.findIngredientWrapperTags(classNames)
        if (ingredientWrapperTags.length > 0) {
            // - look for list of regular tags
            const contents = await HtmlReader.readInsideElement(html, (name, classList) => classList.indexOf(ingredientWrapperTags[0]) > -1)

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
}
