import axios from "axios"
import {backendUrl, Header} from "../utils/urls"
import {Inspection} from "../utils/types";

const url = "inspections/"

const getAllInspection = () => {
    return axios.get(backendUrl + url, Header()); // GET /inspections
}

const updateInspection = (item: Inspection) => {
    return axios.put(backendUrl + url + item.id, item, Header()); // PUT /inspections/:id
}

const addInspection = (item: Inspection) => {
    return axios.post(backendUrl + url, item, Header()); // POST /inspections
}

const deleteInspection = (id: number | null) => {
    return axios.delete(backendUrl + url + id, Header()); // DELETE /inspections/:id
}

export {
    getAllInspection,
    updateInspection,
    addInspection,
    deleteInspection
}