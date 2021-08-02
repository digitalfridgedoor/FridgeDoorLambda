import { HtmlSection } from "./HtmlSection"
import { IngredientSection } from "./Models"

export class IngredientWrapperLocator {

    static find(section: HtmlSection): IngredientSection[] {

        // Assume if children are ingredient list, then this is a wrapper
        const childrenIngredientList = section.children.reduce((acc, curr) => {
            const r = this.find(curr)
            acc.push(...r)
            return acc
        }, [])

        if (childrenIngredientList.length > 0) {
            return childrenIngredientList
        }

        const maybeIngredientWrapper = IngredientWrapperLocator.mightBeAnIngredientWrapper(section)

        if (maybeIngredientWrapper) {
            const ingredients = []
            for (let i = 0; i < section.children.length; i++) {
                const text = section.children[i].text
                if (!text) {
                    return []
                }
                ingredients.push(text)
            }

            return [{ ingredients }]
        }

        let found = false
        let result: IngredientSection[] = []
        section.children.forEach(c => {
            if (found) return

            const r = this.find(c)

            if (r.length > 0) {
                found = true
                result = r
            }
        })

        return result
    }

    private static mightBeAnIngredientWrapper(section: HtmlSection): boolean {
        if (section.children.length === 0) {
            return false
        }
        if (section.children.length === 1) {
            // probably not
            return false
        }

        const expectedTag = section.children[0].tag
        for (let i = 0; i < section.children.length; i++) {
            const t = section.children[i]
            if (t.tag !== expectedTag) {
                return false
            }
        }

        return true
    }
}
