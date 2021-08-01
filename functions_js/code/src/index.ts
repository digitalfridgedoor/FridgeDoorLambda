import { getHtml } from './getHtml';

exports.handler = async function (event, context, callback) {
    console.log('Hello, world');

    const message = event.Records[0].Sns.Message;
    const html = await getHtml(message.url)
    console.log(html)
};
