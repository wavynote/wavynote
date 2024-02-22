export type Folders = {
    folder_id: string;
    folder_name: string;
    note_count: number;
}

export async function getFolderList( userId:string ){
    fetch(`/wavynote/v1.0/main/folderlist?id=${userId}`, {
        method:'GET',
        cache: 'no-store',
        headers: {
          'Authorization':'Basic d2F2eW5vdGU6d2F2eTIwMjMwOTE0',
          'Content-Type':'application/json',
        },})
    .then((res) => { res.json(); })
    .then((data) => { return data });
}