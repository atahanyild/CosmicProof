import { generateGenesisProof, generateNextProof } from './GenerateBlockProof';
import * as process from 'process';

async function main() {
  const args = process.argv.slice(2);
  const command = args[0];

  if (command === 'genesis') {
    const blockHash = args[1];
    if (!blockHash) {
      console.error('Please provide a block hash for the genesis proof.');
      process.exit(1);
    }
    await generateGenesisProof(blockHash);
  } else if (command === 'next') {
    const newBlockHash = args[1];
    const earlierProofJSON = args[2];
    if (!newBlockHash || !earlierProofJSON) {
      console.error(
        'Please provide a new block hash and earlier proof for the next block proof.'
      );
      process.exit(1);
    }
    await generateNextProof(newBlockHash, earlierProofJSON);
  } else {
    console.error('Unknown command. Use "genesis" or "next".');
    process.exit(1);
  }
}

main().catch((err) => {
  console.error('Error:', err);
  process.exit(1);
});
