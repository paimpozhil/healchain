<template>
	<main-layout>
	  <v-flex xs12>
		<v-card class="card--flex-toolbar">
		  <v-toolbar card class="light-blue">
			<v-toolbar-title class="white--text">Transactions</v-toolbar-title>
		  </v-toolbar>
	  
		  <v-data-table :headers="headers" :items="items" hide-actions class="elevation-1">
			<template slot="items" scope="props">
			  <td class="text-xs-left">{{ props.item.to }}</td>
			  <td class="text-xs-left">{{ props.item.from }}</td>
			  <td class="text-xs-left">{{ props.item.amount }}</td>
			</template>
		  </v-data-table>
	  
		</v-card>
	  </v-flex>
	</main-layout>
</template>
<script>
	
import MainLayout from '../layouts/Main.vue'
import { CONTRACT } from '../contract'

export default {
  data () {
    return {
      headers: [
        { align: 'left', text: 'To', value: 'to' },
        { align: 'left', text: 'From', value: 'from' },
        { align: 'left', text: 'Amount', value: 'amount' },
        { align: 'left', text: 'Date', value: '' },
        { align: 'left', text: 'Notes', value: '' },
      ],
      items: []
    }
  },
  
  components: {
      MainLayout
  },

  mounted () {

    CONTRACT.Transfer({}, { fromBlock: 0, toBlock: 'pending' }, (err, res) => {
      if (res.args.to === CONTRACT._eth.coinbase || res.args.from === CONTRACT._eth.coinbase) {
        this.items.push({
          to: res.args.to,
          from: res.args.from,
          amount: web3.fromWei(res.args.value).toNumber()
        })
      }
    })
  }
}
</script>

