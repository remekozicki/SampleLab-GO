import axios from "axios";
import {backendUrl, Header} from "../utils/urls";
import {Indication} from "../utils/types";

const url = 'indications/'

const getIndicationsForSample = (sampleId: string | undefined) => {
    if (!sampleId) return null;
    return axios.get(backendUrl + `samples/${sampleId}/indications`, Header()); // GET /samples/:id/indications
}

const getIndicationById = (indicationId: string | undefined) => {
    if (!indicationId) return null;
    return axios.get(backendUrl + url + indicationId, Header()); // GET /indications/:id
}

const getAllIndications = () => {
    return axios.get(backendUrl + url, Header()); // GET /indications
}

const updateIndication = (item: Indication) => {
    return axios.put(backendUrl + url + item.id, item, Header()); // PUT /indications/:id
}

const addIndication = (item: Indication) => {
    return axios.post(backendUrl + url, item, Header()); // POST /indications
}

const deleteIndication = (id: number | null) => {
    return axios.delete(backendUrl + url + id, Header()); // DELETE /indications/:id
}

export {
    getIndicationsForSample,
    getIndicationById,
    getAllIndications,
    updateIndication,
    addIndication,
    deleteIndication
}
