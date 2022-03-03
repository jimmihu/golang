import React from 'react'
import {
  CCard,
  CCardBody,
  CCol,
  CContainer,
  CFormInput,
  CRow,
  CFormLabel
} from '@coreui/react'
import axios from '../axios'
import { useParams } from 'react-router-dom/cjs/react-router-dom.min'

const Gouserdetail = () => {
    const {id} = useParams()
    const [user, setUser] = React.useState([])
    const [risk, setRisk] = React.useState([])
        axios
          .get(`/user/`+id)
          .then((res) => {
            if (res.data) {
              localStorage.setItem("users", JSON.stringify(res.data.data));
              setUser(res.data.data)
              setRisk(res.data.data.Risk_profile)
            }
          })
          .catch((err) => {
            console.log(err);
          });

  return (
    <div className="bg-light min-vh-100 d-flex flex-row align-items-center">
      <CContainer>
        <CRow className="justify-content-center">
          <CCol md={9} lg={7} xl={6}>
            <CCard className="mx-4">
              <CCardBody className="p-4">
              <h1>{user.name} Detail</h1>
              <CRow className="mb-3">
                <CFormLabel htmlFor="id" className="col-sm-12 col-form-label">ID</CFormLabel>
                <CCol sm={8}>
                  <CFormInput type="text" id="id" defaultValue={user.id} readOnly />
                </CCol>
              </CRow>
              <CRow className="mb-3">
                <CFormLabel htmlFor="name" className="col-sm-12 col-form-label">Name</CFormLabel>
                <CCol sm={8}>
                  <CFormInput type="text" id="name" defaultValue={user.name} readOnly />
                </CCol>
              </CRow>
              <CRow className="mb-3">
                <CFormLabel htmlFor="age" className="col-sm-12 col-form-label">Age</CFormLabel>
                <CCol sm={8}>
                  <CFormInput type="text" id="age" defaultValue={user.age} readOnly />
                </CCol>
              </CRow>
              <CRow className="mb-3">
                <CFormLabel htmlFor="mm" className="col-sm-12 col-form-label">MM Percent(%)</CFormLabel>
                <CCol sm={8}>
                  <CFormInput type="text" id="mm" defaultValue={risk.mm_percent} readOnly />
                </CCol>
              </CRow>
              <CRow className="mb-3">
                <CFormLabel htmlFor="bond" className="col-sm-12 col-form-label">Bond Percent(%)</CFormLabel>
                <CCol sm={8}>
                  <CFormInput type="text" id="bond" defaultValue={risk.bond_percent} readOnly />
                </CCol>
              </CRow>
              <CRow className="mb-3">
                <CFormLabel htmlFor="stock" className="col-sm-12 col-form-label">Stock Percent(%)</CFormLabel>
                <CCol sm={8}>
                  <CFormInput type="text" id="stock" defaultValue={risk.stock_percent} readOnly />
                </CCol>
              </CRow>
              </CCardBody>
            </CCard>
          </CCol>
        </CRow>
      </CContainer>
    </div>
  )
}

export default Gouserdetail