const Bot = require('./lib/Bot')
const SOFA = require('sofa-js')
const Fiat = require('./lib/Fiat')
const http = require('http');

let bot = new Bot()

// ROUTING

bot.onEvent = function(session, message) {
  switch (message.type) {
    case 'Init':
      welcome(session)
      break
    case 'Message':
      onMessage(session, message)
      break
    case 'Command':
      onCommand(session, message)
      break
    case 'Payment':
      onPayment(session, message)
      break
    case 'PaymentRequest':
      welcome(session)
      break
  }
}

function onMessage(session, message) {
  welcome(session)
}

function onCommand(session, command) {
  switch (command.content.value) {
    case 'sos':
      sos(session)
      break
    case 'mywallet':
      count(session)
      break
    case 'brequest':
      brequest(session)
      break
    }
}

function onPayment(session, message) {
  if (message.fromAddress == session.config.paymentAddress) {
    // handle payments sent by the bot
    if (message.status == 'confirmed') {
      // perform special action once the payment has been confirmed
      // on the network
    } else if (message.status == 'error') {
      // oops, something went wrong with a payment we tried to send!
    }
  } else {
    // handle payments sent to the bot
    if (message.status == 'unconfirmed') {
      // payment has been sent to the ethereum network, but is not yet confirmed
      sendMessage(session, `Thanks for the payment! 🙏`);
    } else if (message.status == 'confirmed') {
      // handle when the payment is actually confirmed!
    } else if (message.status == 'error') {
      sendMessage(session, `There was an error with your payment!🚫`);
    }
  }
}

// STATES

function welcome(session) {
  sendMessage(session, `HealChain! What can we do for you?`)
}

function sos(session) {
  sendMessage(session, `Your Friends/Family has been notified with your location for your medical emergency. HealChain Network will look into emergency services.`)
//More sos logic will follow
var options = {
  host: '172.96.13.85',
  port: 1234,
  path: '/profile/' + session.get('paymentAddress')
};

http.get(options, function(res) {
  res.setEncoding('utf8');
  res.on('data', function (chunk) {
    sendMessage(session, chunk);
  });
}).on('error', function(e) {
  console.log("Got error: " + e.message);
});

}

// example of how to store state on each user
function count(session) {
  let count = (session.get('count') || 0) + 1
  session.set('count', count)
  sendMessage(session, `${count}`)
}

function brequest(session) {
  sendMessage(session, `What group & where?`)

//future logic for blood request will be added here

}

function paybill(session) {
sendMessage(session, `TBC : when ERC20 tokens are suppered in Toshi`)
//This functionality will be enabled when Toshi allows ERC20 

}

// HELPERS

function sendMessage(session, message) {
  let controls = [
    {type: 'button', label: 'SOS', value: 'sos'},
    {type: 'button', label: 'MyHealChain', action: "Webview::http://172.96.13.85"},
    {type: 'button', label: 'Blood Request', value: 'brequest'},
    {type: 'button', label: 'Pay Bill', value: 'paybill'}
  ]
  session.reply(SOFA.Message({
    body: message,
    controls: controls,
    showKeyboard: false,
  }))
}
