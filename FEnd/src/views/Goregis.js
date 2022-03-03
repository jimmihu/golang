import React, { useState } from 'react'
import {
  CButton,
  CCard,
  CCardBody,
  CCol,
  CContainer,
  CForm,
  CFormInput,
  CInputGroup,
  CInputGroupText,
  CRow,
} from '@coreui/react'
import CIcon from '@coreui/icons-react'
import { cilCalendar, cilLockLocked, cilUser } from '@coreui/icons'
import axios from '../axios'

const Goregis = () => {
  const [name,setName]=useState("")
  const [password,setPassword]=useState("")
  const [age,setAge]=useState(0)

  const handleRegis = (e) => {
    e.preventDefault()
    console.log("name :", name)
    console.log("age :", age)
    console.log("password :", password)
    axios
      .post(`/user`, { 
            name: name, 
            age: parseInt(age), 
            password: password})
      .then((res) => {
        console.log(res.data)
        if (res.data) {
          localStorage.setItem("user", JSON.stringify(res.data.data));
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    <div className="bg-light min-vh-100 d-flex flex-row align-items-center">
      <CContainer>
        <CRow className="justify-content-center">
          <CCol md={9} lg={7} xl={6}>
            <CCard className="mx-4">
              <CCardBody className="p-4">
                <CForm onSubmit={handleRegis}>
                  <h1>Create User</h1>
                  <p className="text-medium-emphasis">Create your account</p>
                  <CInputGroup className="mb-3">
                    <CInputGroupText>
                      <CIcon icon={cilUser} />
                    </CInputGroupText>
                    <CFormInput placeholder="Username" value={name} onChange={(e)=>setName(e.target.value)} />
                  </CInputGroup>
                  <CInputGroup className="mb-3">
                    <CInputGroupText>
                      <CIcon icon={cilCalendar} />
                    </CInputGroupText>
                    <CFormInput type="number" placeholder="Age" value={age} onChange={(e)=>setAge(e.target.value)} />
                  </CInputGroup>
                  <CInputGroup className="mb-3">
                    <CInputGroupText>
                      <CIcon icon={cilLockLocked} />
                    </CInputGroupText>
                    <CFormInput
                      value={password} onChange={(e)=>setPassword(e.target.value)}
                      type="password"
                      placeholder="Password"
                    />
                  </CInputGroup>
                  <div className="d-grid">
                    <CButton color="success">Create</CButton>
                  </div>
                </CForm>
              </CCardBody>
            </CCard>
          </CCol>
        </CRow>
      </CContainer>
    </div>
  )
}

export default Goregis
