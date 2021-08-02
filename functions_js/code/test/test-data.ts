import { SNSEvent } from 'aws-lambda';
import * as faker from 'faker';

export function createSNSMessage(message: string): SNSEvent {
    return {
        Records: [{
            EventVersion: faker.random.word(),
            EventSubscriptionArn: faker.random.word(),
            EventSource: 'SNS',
            Sns: {
                SignatureVersion: faker.random.word(),
                Timestamp: faker.random.word(),
                Signature: faker.random.word(),
                SigningCertUrl: faker.random.word(),
                MessageId: faker.random.word(),
                Message: message,
                MessageAttributes: {},
                Type: faker.random.word(),
                UnsubscribeUrl: faker.random.word(),
                TopicArn: faker.random.word(),
                Subject: faker.random.word(),
            }
        }]
    }
};