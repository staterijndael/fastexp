import server_url from "../server_url";

export class TagList {
  static get_command_tags_url = "/private/get_command_tags";
  static get_user_tags_url = "/private/get_user_tags";
  static set_user_tags_url = "/private/set_user_tags";

  static async get_command_tags(command_id) {
    // let myHeaders = new Headers();
    // myHeaders.append("Content-Type", "application/json");

    // let raw = JSON.stringify({
    //   command_id: command_id,
    // });

    // let requestOptions = {
    //   method: "POST",
    //   headers: myHeaders,
    //   body: raw,
    // };

    let tag_list = {
      err: null,
      tags: null,
    };

    // try {
    //   let res = await fetch(
    //     server_url + TagList.get_command_tags_url,
    //     requestOptions
    //   );
    //   if (res.ok) {
    //     res = await res.json();
    //     tag_list.tags = res.tags;
    //   } else {
    //     tag_list.err = true;
    //   }
    // } catch (err) {
    //   tag_list.err = err;
    // }

    tag_list = {
        err: null,
        tags: ['Docker', 'Go RabbitMq', 'uber-go', 'Aliases web-pack', 'babel', 'SOLID']
    }

    return tag_list;
  }

  static async get_user_tags(user_id) {
    let myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    let raw = JSON.stringify({
      user_id: user_id,
    });

    let requestOptions = {
      method: "POST",
      headers: myHeaders,
      body: raw,
    };

    let tag_list = {
      err: null,
      tags: null,
    };

    try {
      let res = await fetch(
        server_url + TagList.get_user_tags_url,
        requestOptions
      );
      if (res.ok) {
        res = await res.json();
        tag_list.tags = res.tags;
      } else {
        tag_list.err = true;
      }
    } catch (err) {
      tag_list.err = err;
    }

    return tag_list;
  }

  static async set_user_tags(list_of_tags) {
    let myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    let raw = JSON.stringify({
      tag_list: list_of_tags,
    });

    let requestOptions = {
      method: "POST",
      headers: myHeaders,
      body: raw,
    };

    let isOk = true;
    let res = await fetch(
      server_url + this.set_person_tags_url,
      requestOptions
    );
    if (!res.ok) isOk = false;

    return isOk;
  }
}
