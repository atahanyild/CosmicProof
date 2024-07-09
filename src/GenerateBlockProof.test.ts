import { GenerateBlockProof } from './GenerateBlockProof';
import { Field } from 'o1js';

let proofsEnabled = false;
describe('GenerateBlockProof', () => {
  beforeAll(async () => {
    if (proofsEnabled) await GenerateBlockProof.compile();
  });

  it('correctly updates the state on the `Fibonacci` smart contract', async () => {
    let proof = await GenerateBlockProof.generateGenesisBlockProof(
      Field('676876768')
    );
    console.log(
      'current proof of block: ',
      proof.proof,
      proof.publicOutput.toJSON(),
      proof.publicInput.toJSON()
    );
    let hash2 = '123123';
    let proof2 = await GenerateBlockProof.generateNextBlockProof(
      Field(hash2),
      proof
    );
    console.log(
      'current proof of block: ',
      proof2.proof,
      proof2.publicOutput.toJSON(),
      proof2.publicInput.toJSON()
    );

    const verify2 = await GenerateBlockProof.verify(proof2);
    console.log('verify2: ', verify2);

    let hash3 = '456456';
    let proof3 = await GenerateBlockProof.generateNextBlockProof(
      Field(hash3),
      proof2
    );
    console.log(
      'current proof of block: ',
      proof3.proof,
      proof3.publicOutput.toString(),
      proof3.publicInput.toString()
    );

    const verify3 = await GenerateBlockProof.verify(proof3);
    console.log('verify3: ', verify3);
  });
});
