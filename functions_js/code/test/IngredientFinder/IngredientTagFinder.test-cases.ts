import * as fs from 'fs'
import { IngredientSection } from '../../src/IngredientFinder/Models';
import 'mocha';

export interface ITestCase {
    name: string;
    html: string;
    expected: IngredientSection[];
}

function readTestCase(fileName: string): string {
    return fs.readFileSync('test/TestData/' + fileName + '.html', { encoding: 'utf-8' })
}

const testCases: { file: string, expected: IngredientSection[] }[] = [{
    file: 'jamie_cauli',
    expected: [{
        ingredients: ['1 small head of cauliflower , (600g)',
            '2 heaped teaspoons rose harissa',
            '2 heaped tablespoons natural yoghurt',
            '1 pomegranate',
            '1 tablespoon dukkah']
    }]
},
{
    file: 'chicken_shwarma',
    expected: [{
        ingredients: ['2lb /1 kg chicken thigh fillets, skinless and boneless (Note 3)']
    }, {
        ingredients: [
            '1 large garlic clove, minced (or 2 small cloves)',
            '1 tbsp ground coriander',
            '1 tbsp ground cumin',
            '1 tbsp ground cardamon',
            '1 tsp ground cayenne pepper (reduce to 1/2 tsp to make it not spicy)',
            '2 tsp smoked paprika',
            '2 tsp salt',
            'Black pepper',
            '2 tbsp lemon juice',
            '3 tbsp olive oil',]
    }, {
        'ingredients': [
            '1 cup Greek yoghurt',
            '1 clove garlic, crushed',
            '1 tsp cumin',
            'Squeeze of lemon juice',
            'Salt and pepper'
        ]
    },
    {
        'ingredients': [
            '6 flatbreads (Lebanese or pita bread orhomemade soft flatbreads)',
            'Sliced lettuce (cos or iceberg)',
            'Tomato slices'
        ]
    }]
},
// {
//     file: 'shashuska_traybake',
//     expected: [{
//         ingredients: ['2 red peppers, thinly sliced',
//             '2 yellow peppers, thinly sliced',
//             '1 red onion, thickly sliced',
//             '1 aubergine, sliced into 15mm/⅝in strips',
//             '800g/1lb12oz vine tomatoes, quartered',
//             '1 red chilli, roughly chopped',
//             '1 tsp ground cumin',
//             '1 tsp smoked paprika',
//             '2 tbsp olive oil',
//             '4 free-range eggs',
//             'sea salt and freshly ground black pepper',
//             'handful fresh coriander, roughly chopped, to serve',
//             'flatbreads or pitta breads, warmed, to serve',]
//     }]
// },
{
    file: 'thai_green_curry',
    expected: [{
        ingredients: ['2 tsp. vegetable oil',
            '400 g skinless boneless chicken breast, cut into bite-size pieces',
            '3 spring onions, sliced diagonally',
            '1/2 red pepper, thinly sliced into strips',
            '1 garlic clove, crushed',
            'Small handful Thai basil leaves, ripped',
            '50 g Thai green curry paste (about 4 tbsp)',
            '400 ml chicken stock',
            '160 ml can coconut cream',
            '1 tsp. fish sauce',
            '175 g baby sweetcorn, sliced in thirds diagonally',
            '225 g tin bamboo shoots, thoroughly drained',
            '100 g mangetout, sliced in half diagonally',
            'Small handful Thai basil leaves, ripped, and lime wedges, to serve',]
    }]
},
    // {
    //     file: 'pot_curry',
    //     expected: [{
    //         ingredients: ['2 tsp. vegetable oil',
    //         ]
    //     }]
    // }
]

export const TestCases: ITestCase[] = testCases.map(({ file, expected }) => {
    const html = readTestCase(file)
    return { name: file, html, expected };
})
