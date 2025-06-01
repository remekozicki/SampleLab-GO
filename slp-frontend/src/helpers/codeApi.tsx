import axios from "axios"
import {backendUrl, Header} from "../utils/urls"
import {Code} from "../utils/types";

const url = 'codes/'

const getAllCodes = () => {
    return axios.get(backendUrl + url, Header()) // GET /codes
}

const updateCode = (item: Code) => {
    return axios.put(backendUrl + url + item.id, item, Header()) // PUT /codes/:id
}

const addCode = (item: Code) => {
    return axios.post(backendUrl + url, item, Header()) // POST /codes
}

const deleteCode = (id: string | null) => {
    return axios.delete(backendUrl + url + id, Header()) // DELETE /codes/:id
}

export {
    getAllCodes,
    updateCode,
    addCode,
    deleteCode
}