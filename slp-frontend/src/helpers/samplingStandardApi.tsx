import axios from "axios"
import {SamplingStandards} from "../utils/types";
import {backendUrl, Header} from "../utils/urls"

const url = "sampling-standards/"

const getAllSamplingStandard = () => {
    return axios.get(backendUrl + url, Header()); // GET /sampling-standards
}

const updateSamplingStandard = (item: SamplingStandards) => {
    return axios.put(backendUrl + url + item.id, item, Header()); // PUT /sampling-standards/:id
}

const addSamplingStandard = (item: SamplingStandards) => {
    return axios.post(backendUrl + url, item, Header()); // POST /sampling-standards
}

const deleteSamplingStandard = (id: number | null) => {
    return axios.delete(backendUrl + url + id, Header()); // DELETE /sampling-standards/:id
}

export {
    getAllSamplingStandard,
    updateSamplingStandard,
    addSamplingStandard,
    deleteSamplingStandard
}