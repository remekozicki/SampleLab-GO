import axios from "axios"
import {backendUrl, Header} from "../utils/urls"
import {ProductGroupSave} from "../utils/types";

const url = "product-groups/"

const getAllGroup = () => {
    return axios.get(backendUrl + url, Header()) // GET /product-groups
}

const updateGroup = (item: ProductGroupSave) => {
    return axios.put(backendUrl + url + item.id, item, Header()) // PUT /product-groups/:id
}

const addGroup = (item: ProductGroupSave) => {
    return axios.post(backendUrl + url, item, Header()) // POST /product-groups
}

const deleteGroup = (id: number | null) => {
    return axios.delete(backendUrl + url + id, Header()) // DELETE /product-groups/:id
}

export {
    getAllGroup,
    updateGroup,
    addGroup,
    deleteGroup
}

