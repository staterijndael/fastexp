const server_url = require('../server_url')

class Card {
    static get_card_url = '/private/get_card'

    static async get_card(theme_id) {
        let myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/json');

        let raw = JSON.stringify({
            'theme_id': theme_id
        });

        let requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
        };

        let card = {
            'err': null,
            'title': null,
            'content': null,
        };

        try {
            let res = await fetch(server_url + Theme.get_card_url, requestOptions)
            if (res.ok) {
                res = await res.json()
                card.title = res.title
                card.content = res.content
            } else {
                card.err = true
            }
        } catch(err) {
            card.err = err
        }
        
        return card
    }
}

module.exports = Card;
