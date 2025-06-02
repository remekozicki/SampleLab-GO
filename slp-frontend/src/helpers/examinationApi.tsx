import axios from "axios";
import {backendUrl, Header} from "../utils/urls";
import {Examination} from "../utils/types";

const url = 'examination/'

const getExaminationsForSample = (sampleId: string | undefined) => {
    if (sampleId !== undefined) {
        return axios.get(backendUrl + url + sampleId, Header())
    }
    return null;
}

const getExaminationById = (examinationId: string | undefined) => {
    if (examinationId !== undefined) {
        return axios.get(backendUrl + url + `${examinationId}`, Header());
    }
    return null;
}

const updateExamination = (updatedExamination: Examination) => {
    return axios.put(backendUrl + url, updatedExamination, Header());
}

const addExamination = (examination: Examination) => {
    return axios.post(backendUrl + url, examination, Header());
}

const deleteExamination = (examinationId: number) => {
    return axios.delete(backendUrl + url + examinationId, Header());
}

export {
    getExaminationsForSample,
    getExaminationById,
    updateExamination,
    addExamination,
    deleteExamination,
}
