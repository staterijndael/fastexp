import { server_url } from "../server_url";

export class User {
    static whoami_url = '/private/whoami'

    static async whoami() {
        let requestOptions = {
            method: 'GET',
        };
    
        let user = {
            err: null,
            id: null,
            email: null,
            tags: null,
        }

        try {
            let res = await fetch(server_url + User.whoami, requestOptions)
            if (res.ok) {
                res = await res.json()
                user.id = res.id
                user.email = res.email
                user.tags = res.tags
            } else {
                user.err = true
            }
        } catch(err) {
            user.err = err
        } 

        return user
    }
}
