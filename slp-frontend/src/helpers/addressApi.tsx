import axios from "axios"
import { backendUrl } from "../utils/urls"
import { Header } from "../utils/urls"
const url = 'address/'

const getAllAddresses = () => {
    return axios.get(backendUrl + url, Header())
}

export {
    getAllAddresses
}