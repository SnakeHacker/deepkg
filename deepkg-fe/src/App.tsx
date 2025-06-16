import { Routes, Route, HashRouter } from "react-router-dom";
import './App.css'
import LoginPage from "./pages/login";
import { AuthRoute } from './auth';
import NotFoundPage from "./pages/404/NotFondPage";
import HomePage from "./pages/home";
import UserListPage from "./pages/user";
import OrganizationListPage from "./pages/orginization";
function App() {

  return (
    <>
      <HashRouter>
        <Routes>

            <Route path="/login" element={<LoginPage />} />
            <Route path="/" element={<AuthRoute><HomePage /></AuthRoute>}>

                <Route index element={<UserListPage />} />
                <Route path="org" element={<OrganizationListPage />} />
                <Route path="user" element={<UserListPage />} />
            </Route>
            <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </HashRouter>
    </>
  )
}

export default App
