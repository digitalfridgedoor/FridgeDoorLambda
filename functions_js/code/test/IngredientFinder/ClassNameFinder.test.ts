import { ClassNameFinder } from '../../src/IngredientFinder/ClassNameFinder';
import { expect } from 'chai';
import 'mocha';

describe('ClassNameFinder', () => {

  describe('getAllClassNameFinder', async () => {
    const html = `<div class="wrapper">
      <div class="ingredient" />
      <div class="ingredient" />
      <div class="ingredient" />
      <div>`

    const results = await ClassNameFinder.getAllClassNames(html)

    it('Should find 2 class names', async () => {
      expect(Object.keys(results).length).to.equal(2);
    });
    it('Should find 1 wrapper instance', async () => {
      expect(results['wrapper']).to.equal(1);
    });
    it('Should find 3 ingredient instances', async () => {
      expect(results['ingredient']).to.equal(3);
    });
  });
});
