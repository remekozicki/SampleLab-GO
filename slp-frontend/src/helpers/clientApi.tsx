import axios from "axios"
import {backendUrl, Header} from "../utils/urls"
import {Client} from "../utils/types";

const url = "clients/"

const getAllClients = () => {
    return axios.get(backendUrl + url, Header()); // GET /clients
}

const updateClient = (item: Client) => {
    return axios.put(backendUrl + url + item.id, item, Header()); // PUT /clients/:id
}

const addClient = (item: Client) => {
    return axios.post(backendUrl + url, item, Header()); // POST /clients
}

const deleteClient = (id: number | null) => {
    return axios.delete(backendUrl + url + id, Header()); // DELETE /clients/:id
}

export {
    getAllClients,
    updateClient,
    addClient,
    deleteClient
}