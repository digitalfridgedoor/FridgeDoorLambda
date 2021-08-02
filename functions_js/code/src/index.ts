import { SNSEvent } from "aws-lambda";
import { getHtml } from './getHtml';

export const handler = async (
    event: SNSEvent
): Promise<void> => {
    console.log('Hello, world');

    const message = event.Records[0].Sns.Message;
    console.log(event.Records[0].Sns);
    const messageContents = JSON.parse(message);

    console.log('getting html from ' + messageContents.url);
    const html = await getHtml(messageContents.url)
    console.log(html)
};
