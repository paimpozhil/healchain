import config from './config'

const CONTRACT = web3.eth.contract(config.contractAbi).at(config.contractAddress, (err, ctr) => {
  return ctr
})

export { CONTRACT }