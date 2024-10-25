import PocketBase from 'npm:pocketbase';

const pb = new PocketBase('http://192.168.1.253:8090');

async function controlLight(roomName: string, lightName: string, durationms: string) {
    await pb.send(`/rooms/${roomName}/lights/${lightName}`, { duration: durationms });
}

await controlLight("MakerFaire", "LuceTavolo", "1000");