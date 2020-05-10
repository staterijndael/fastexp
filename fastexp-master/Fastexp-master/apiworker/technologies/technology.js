const server_url = require('../server_url')

class TechnologyList {
    static get_technology_list_url = '/private/get_technology_list'

    static async get_technology_list(user_id) {
        let myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/json');
        
        let raw = JSON.stringify({
            'user_id': user_id
        });

        let requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
        };

        let technology_list = {
            'err': null,
            'techs': null,
        };

        try {
            let res = await fetch(server_url + TechnologyList.get_technology_list_url, requestOptions)
            if (res.ok) {
                res = await res.json()
                technology_list.techs = res.techs
            } else {
                technology_list.err = true
            }
        } catch(err) {
            technology_list.err = err
        }
        
        return technology_list
    }
}

module.exports = TechnologyList;
