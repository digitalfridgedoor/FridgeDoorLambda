import { Parser } from 'htmlparser2';

export class ClassNameFinder {

    static getAllClassNames(pageHtml: string): Promise<{ [key: string]: number }> {
        return new Promise((resolve, reject) => {
            const tags: { [key: string]: number } = {}

            var parser = new Parser({
                onopentag: (name, attribs) => {
                    if (!!attribs.class) {
                        const classList = attribs.class.split(' ')
                        classList.forEach(c => {
                            if (c) {
                                if (typeof tags[c] === 'undefined') {
                                    tags[c] = 0
                                }
                                tags[c]++
                            }
                        })
                    }
                },
                onend: () => {
                    resolve(tags);
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
