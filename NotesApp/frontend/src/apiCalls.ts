const baseUrl = "http://localhost:1000/notes";

interface NotePayload {
    header: string,
    content: string,
    user?: string
}
const setOptions = function (token: string, httpVerb: string, payload?: any) {
    return {
        method: httpVerb,
        headers: new Headers({
            'Authorization': 'Basic ' + token,
            'Content-Type': 'application/x-www-form-urlencoded',
        }),
        body: JSON.stringify(payload)
    } 
}


export const getNote = async function (token: string, id: string) {
    const requestOptions = setOptions(token, "GET");
    const response = await fetch(`${baseUrl}/${id}`,requestOptions)
    const { note } = (await response.json());
    return note;
}

export const getNotes = async function (token: string) {
    const requestOptions = setOptions(token, "GET");
    const response = await fetch(baseUrl, requestOptions)
    const { notes } = (await response.json());
    return notes;

}

export const saveNote = async function (token: string, { header, content }: NotePayload) {
    const requestOptions = setOptions(token, "POST", { header, content });
    const response = await fetch(baseUrl, requestOptions)
    return response
}

export const deleteNote =  async function (token:string,noteId:string){
    const requestOptions =  setOptions(token,"DELETE");
    const response = await fetch(`${baseUrl}/${noteId}`,requestOptions)
    return response
}

export const shareNote = async function (token:string ,noteId:string,email:string) {
    const requestOptions =  setOptions(token,"POST",{email,
        noteID:noteId
    });
    const response = await fetch(`${baseUrl}/share`,requestOptions)
    return response
    
}

export const updateNote = async function (token: string, { header, content, noteID }:any) {
    const requestOptions = setOptions(token, "POST", { content,header,noteID });
    const response = await fetch(baseUrl, requestOptions)
    return response
}

