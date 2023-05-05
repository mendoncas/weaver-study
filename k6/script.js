import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
    let url = "http://localhost:12345"

    let bookData = {
        title: "",
        description: "",
        author: ""
    }

    for (key in bookData.keys)
        bookData[key] = (Math.random() + 1).toString(36).substring(7);

    http.post(url, JSON.stringify(bookData), {
        headers: {'Content-Type': 'application/json'}
    })
    sleep(1);
}
