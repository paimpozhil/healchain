<template>
  <v-app dark>
    <v-toolbar fixed>
      <v-toolbar-title v-text="title"></v-toolbar-title>
    </v-toolbar>
    <main>
      <v-alert error value="true" v-if="!metamask">
        This application requires the MetaMask Chrome Plugin
      </v-alert>
      <template v-if="metamask">
        <v-container>
          <v-layout row class="pb-5">
            <wallet-info></wallet-info>
            <send-tokens></send-tokens>
          </v-layout>
          <v-layout row class="pb-5">
            <wallet-info-user></wallet-info-user>
          </v-layout>
          <v-layout row class="pb-5">
            <donate-blood></donate-blood>
          </v-layout>
          <v-layout row class="pb-5">
            <transactions></transactions>
          </v-layout>
        </v-container>
      </template>
    </main>
    <v-footer fixed>
      <span>&copy; 2017</span>
    </v-footer>
  </v-app>
</template>

<script>
import WalletInfo from './components/WalletInfo.vue'
import DonateBlood from './components/DonateBlood.vue'
import WalletInfoUser from './components/WalletInfoUser.vue'
import SendTokens from './components/SendTokens.vue'
import Transactions from './components/Transactions.vue'

export default {

  components: {
    WalletInfo,
    WalletInfoUser,
    SendTokens,
    DonateBlood,
    Transactions
  },

  data () {
    return {
      title: 'HealChain'
    }
  },

  computed: {
    metamask () {
      return window['web3'] !== undefined
    }
  }

}
</script>

<style lang="stylus">
  @import './stylus/main'
</style>
