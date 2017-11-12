<template>
  <div class="container">
      <nav class="contact">
        <v-link href="/">Home</v-link>
        <v-link href="/walletinfouser">walletinfouser</v-link>
        <v-link href="/donateblood">donateblood</v-link>
        <v-link href="/sendtoken">sendtoken</v-link>
        <v-link href="/transactions">transactions</v-link>
      </nav>

    <slot></slot>
  </div>
</template>

<script>
  import VLink from '../components/VLink.vue'

  export default {
    components: {
      VLink
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

</style>
