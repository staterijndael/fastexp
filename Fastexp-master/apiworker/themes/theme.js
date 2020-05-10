import server_url from '../server_url';

class Theme {
    static get_themes_url = '/private/get_themes'

    static async get_themes(technology_id) {
        let myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/json');
        
        let raw = JSON.stringify({
            'technology_id': technology_id
        });

        let requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
        };

        let themes = {
            'err': null,
            'themes': null
        };

        try {
            let res = await fetch(server_url + Theme.get_themes_url, requestOptions)
            if (res.ok) {
                res = await res.json()
                themes.themes = res.themes
            } else {
                themes.err = true
            }
        } catch(err) {
            themes.err = err
        }
        
        return themes
    }
}

export default Theme;
