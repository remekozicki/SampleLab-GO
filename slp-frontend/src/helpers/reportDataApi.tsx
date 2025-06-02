import axios from "axios"
import { backendUrl, Header } from "../utils/urls"
import { ReportData } from "../utils/types"

const url = "report-data/"

const getAllReportData = () => {
    return axios.get(backendUrl + url, Header())
}

const addReportData = (reportData:ReportData) => {
    return axios.post(backendUrl + url, reportData, Header())
}

const getReportDataBySampleId = (sampleId:string | undefined) => {
    return axios.get(backendUrl + url + `sample/${sampleId}`, Header())
}

export {
    getAllReportData,
    addReportData,
    getReportDataBySampleId
}