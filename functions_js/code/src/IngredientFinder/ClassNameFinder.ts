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
                    // stack.push({ name, attribs, text: '', selectNode })
                },
                ontext: function (text) {
                    // console.log('text', text)
                    // if (stack.length > 0) {
                    //     text = text.split(/\s+/g).filter(x => !!x).join(' ')
                    //     stack[stack.length - 1].text += text
                    // }
                },
                onclosetag: () => {
                    // insideCount--;
                    // const last = stack.pop()
                    // if (last.selectNode) {
                    //     results.push(options.value(last.name, last.attribs, last.text));
                    // }
                    // if (stack.length > 0) {
                    //     stack[stack.length - 1].text += last.text
                    // }
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
