import { Field, SelfProof, ZkProgram, Poseidon } from 'o1js';

export const GenerateBlockProof = ZkProgram({
  name: 'generate-block-proof',
  publicInput: Field,
  publicOutput: Field,

  methods: {
    generateGenesisBlockProof: {
      privateInputs: [],

      async method(blockHash: Field) {
        return Poseidon.hash([blockHash]);
      },
    },

    generateNextBlockProof: {
      privateInputs: [SelfProof],
      /**
       *
       * @param newBlockHash
       * @param earlierProof
       * @returns
       */

      async method(newBlockHash: Field, earlierProof: SelfProof<Field, Field>) {
        earlierProof.verify();

        return Poseidon.hash([earlierProof.publicOutput, newBlockHash]);
      },
    },
  },
});

await GenerateBlockProof.compile();

export async function generateGenesisProof(blockHash: string) {
  const proof = await GenerateBlockProof.generateGenesisBlockProof(
    Field(blockHash)
  );
  console.log(JSON.stringify(proof.toJSON()));
  return proof;
}

export async function generateNextProof(
  newBlockHash: string,
  earlierProofJSON: string
) {
  const earlierProof = await SelfProof.fromJSON(JSON.parse(earlierProofJSON));
  const proof = await GenerateBlockProof.generateNextBlockProof(
    Field(newBlockHash),
    earlierProof
  );
  console.log(JSON.stringify(proof.toJSON()));
  return proof;
}
