import { getHtml } from './getHtml';

exports.handler = async function (event, context, callback) {
    console.log('Hello, world');

    const message = event.Records[0].Sns.Message;
    console.log(event.Records[0].Sns);
    console.log(message);
    console.log('getting html from ' + message.url);
    const html = await getHtml(message.url)
    console.log(html)
};
