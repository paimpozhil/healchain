<template>
	
  <div class="container">
		<v-layout wrap justify-space-around align-center>
			<img src="/src/assets/logo.jpg" alt="John">
		</v-layout>
	<!-- Handlebars template -->
	<nav  v-show="showing" class="contact" v-if="isLab === 1"> 
        <v-link href="/">Home</v-link>
        <v-link href="/sendtoken">Send Token</v-link>
        <v-link href="/transactions">History</v-link>
        <v-link href="/donateblood">Donor Rebate</v-link>
        <v-link href="/healreport">Heal Report</v-link>
	</nav>
	<nav v-show="showing" class="contact" v-else> 
        <v-link href="/">Home</v-link>
        <v-link href="/sendtoken">Send Token</v-link>
        <v-link href="/transactions">History</v-link>
	</nav>
    <slot></slot>
  </div>
</template>

<script>
	import VLink from '../components/VLink.vue'
	import { CONTRACT } from '../contract'

  export default {
    components: {
      VLink
    },
    
  data () {
    return {
      isLab: 0,
      showing: true,
    }
  },

  mounted () {
    this.checkIsLab()
    
  },

  methods: {
    checkIsLab () {
      CONTRACT.is_lab((err, bal) => {
		  if(bal){
			 this.isLab =  1;
		  }
      })
      if(window.location.href.indexOf("users") > -1) {			
			this.showing=false;
		} 
    },
    setIsLab () {
      CONTRACT.set_lab(1, (err, bal) => {
        console.log(err);
      })
    },

  }
}

</script>

<style scoped>
  .container {
    max-width: 900px;
    margin: 0 auto;
    padding: 15px 30px;
    background: #f9f7f5;
  }
  nav .active {
		background-color: #e35885;
	}	

	nav{
		display:inline-block;
		margin:60px auto 45px;
		background-color:#5597b4;
		box-shadow:0 1px 1px #ccc;
		border-radius:2px;
	}

	nav a{
		display:inline-block;
		padding: 18px 30px;
		color:#fff !important;
		font-weight:bold;
		font-size:16px;
		text-decoration:none !important;
		line-height:1;
		text-transform: uppercase;
		background-color:transparent;

		-webkit-transition:background-color 0.25s;
		-moz-transition:background-color 0.25s;
		transition:background-color 0.25s;
	}

	nav a:first-child{
		border-radius:2px 0 0 2px;
	}

	nav a:last-child{
		border-radius:0 2px 2px 0;
	}

	select {
		-webkit-appearance: menulist;
		box-sizing: border-box;
		align-items: center;
		white-space: pre;
		-webkit-rtl-ordering: logical;
		color: black;
		background-color: white;
		cursor: default;
		border-width: 1px;
		border-style: solid;
		border-color: initial;
		border-image: initial;
	}
</style>
