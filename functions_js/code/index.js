const { getHtml } = require("./getHtml");

exports.handler = async function (event, context, callback) {
    console.log('Hello, world');

    const message = event.Records[0].Sns.Message;
    try {
        const messageData = JSON.parse(message);

        const html = await getHtml(messageData.url)
        console.log(html)
    } catch {
        console.error('Could not parse message: ' + message)
    }
};
