// 1. Import modules.
import { createPublicClient, http } from "viem";
import { mainnet } from "viem/chains";

// 2. Set up your client with desired chain & transport.
const client = createPublicClient({
  chain: mainnet,
  transport: http(),
});

// 3. Consume an action!
const main = async () => {
  const blockNumber = await client.getBlockNumber();
  console.log("blockNumber", BigInt(blockNumber).toString());
};

main();
