import * as https from 'https';
import { IngredientTagFinder } from './IngredientTagFinder';
import { IngredientSection } from './Models';

export class TagFinder {
    constructor() {
    }

    async getPage(url: string): Promise<IngredientSection[]> {
        const html = await this.httpGet(url)

        return await IngredientTagFinder.find(html)
    }

    private async httpGet(url: string): Promise<string> {
        return new Promise<string>((resolve, reject) => {
            https.get(url, (res => {
                if (res.statusCode === 200) {
                    let data = '';
                    res.on('data', d => data += d)
                    res.on('end', () => resolve(data))
                } else {
                    reject({ statusCode: res.statusCode, reason: 'invalid status' })
                }
            }))
        })
    }
}
