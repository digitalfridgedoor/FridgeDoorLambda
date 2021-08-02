import { SNSEvent } from "aws-lambda";
import { getHtml } from './getHtml';
import { IngredientTagFinder } from "./IngredientFinder/IngredientTagFinder";

export const handler = async (
    event: SNSEvent
): Promise<void> => {
    const message = event.Records[0].Sns.Message;
    // console.log(event.Records[0].Sns);

    const messageContents = JSON.parse(message);
    console.log('getting html from ' + messageContents.url);

    const html = await getHtml(messageContents.url)
    const sections = IngredientTagFinder.find(html);
    console.log(sections)
};
