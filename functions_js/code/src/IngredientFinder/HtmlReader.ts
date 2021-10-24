import { Parser } from 'htmlparser2';
import { HtmlSection } from './HtmlSection';

export class HtmlReader {
    constructor() {
    }

    public static async readContents(html: string): Promise<HtmlSection> {
        const section = await this.readHtmlSection(html);
        return section;
    }

    private static async readHtmlSection(html: string): Promise<HtmlSection> {
        try {
            return await HtmlReader.readInsideElement(html, (name, classList) => name === 'body')
        } catch {
            // try reading whole page
            return await HtmlReader.readInsideElement(html, (name, classList) => true)
        }
    }

    public static readInsideElement(pageHtml: string, isTopElement: (name: string, classList: string[]) => boolean): Promise<HtmlSection> {
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
                    console.log(name, insideCount, attribs.class)
                    const classList = !attribs.class ? [] : attribs.class.split(' ')
                    if (insideCount >= 0) {
                        insideCount++
                        pushSection({ classList, tag: name, children: [] })
                        return
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
                    console.log('close')
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
