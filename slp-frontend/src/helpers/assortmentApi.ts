import axios from "axios"
import {backendUrl, Header} from "../utils/urls"
import {Assortment} from "../utils/types";

const url = "assortment/"

const getAllAssortments = () => {
    return axios.get(backendUrl + url, Header())
}
const updateAssortment = (item: Assortment) => {
    return axios.put(backendUrl + url, item, Header());
}

const addAssortment = (item: Assortment) => {
    return axios.post(backendUrl + url, item, Header());
}

const deleteAssortment = (id: number | null) => {
    return axios.delete(backendUrl + url + id, Header());
}

export {
    getAllAssortments,
    updateAssortment,
    addAssortment,
    deleteAssortment
}

