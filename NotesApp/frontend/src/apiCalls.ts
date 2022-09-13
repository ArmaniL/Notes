const baseUrl = "http://localhost:1000/notes";

interface NotePayload {
    header: string,
    content: string,
    user: string
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



export const getNotes = async function (token: string) {
    const requestOptions = setOptions(token, "GET");
    const response = await fetch(baseUrl, requestOptions)
    const { notes } = (await response.json());
    return notes;

}

export const saveNote = async function (token: string, { header, content, user }: NotePayload) {
    const requestOptions = setOptions(token, "POST", { header, content, user });
    const response = await fetch(baseUrl, requestOptions)
}