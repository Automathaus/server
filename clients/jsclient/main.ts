import PocketBase from 'npm:pocketbase';

const pb = new PocketBase('http://192.168.1.253:8090');

await pb.send("/print-body", {
    method: "POST",
    body: { abc: 123 },
});