import React from 'react'
import {
    CCard,
    CCardBody,
    CCol,
    CRow,
    CTable,
    CTableBody,
    CTableDataCell,
    CTableHead,
    CTableHeaderCell,
    CTableRow,
    CContainer,
    CLink
} from '@coreui/react'
import axios from '../axios'

const Gouserlist = () => {
    const [users, setUser] = React.useState([])
        axios
          .get(`/users`, { params:{ 
                limit:10,
                offset:0
            }})
          .then((res) => {
            if (res.data) {
              localStorage.setItem("users", JSON.stringify(res.data.data));
              setUser(res.data.data)
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
                <h1>User List</h1>
              <CTable>
                <CTableHead>
                    <CTableRow>
                    <CTableHeaderCell scope="col">ID</CTableHeaderCell>
                    <CTableHeaderCell scope="col">Name</CTableHeaderCell>
                    <CTableHeaderCell scope="col">Age</CTableHeaderCell>
                    </CTableRow>
                </CTableHead>
                <CTableBody>
                {users.map((user, index) => (  
                    <CTableRow key={index} >
                    <CTableHeaderCell scope="row">{user.id}</CTableHeaderCell>
                    <CTableDataCell><CLink href={"#/user/"+user.id}>{user.name}</CLink></CTableDataCell>
                    <CTableDataCell>{user.age}</CTableDataCell>
                    </CTableRow>
                ))}  
                </CTableBody>
                </CTable>
              </CCardBody>
            </CCard>
          </CCol>
        </CRow>
      </CContainer>
    </div>
  )
}

export default Gouserlist