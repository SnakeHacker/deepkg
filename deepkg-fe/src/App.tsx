import { Routes, Route, HashRouter } from "react-router-dom";
import './App.css'
import LoginPage from "./pages/login";
import { AuthRoute } from './auth';
import NotFoundPage from "./pages/404/NotFondPage";
import HomePage from "./pages/home";
import UserListPage from "./pages/user";
import OrganizationListPage from "./pages/orginization";
import OntologyPage from "./pages/ontology";
import WorkspacePage from "./pages/workspace";
import TriplePage from "./pages/triple";
import DocumentDirPage from "./pages/document_dir";
import DocumentPage from "./pages/document";
import ExtractTaskPage from "./pages/extract_task";
import OntologyPropPage from "./pages/ontology_prop";
import Dashboard from "./pages/dashboard";
import ExtractTaskResultPage from "./pages/extract_task_result";
import KnowledgeInferPage from "./pages/knowledge_infer";


function App() {
  return (
    <>
      <HashRouter>
        <Routes>

            <Route path="/login" element={<LoginPage />} />
            <Route path="/" element={<AuthRoute><HomePage /></AuthRoute>}>

                <Route index element={<Dashboard />} />
                <Route path="dashboard" element={<Dashboard />} />
                <Route path="document_dir" element={<DocumentDirPage />} />
                <Route path="document" element={<DocumentPage />} />
                <Route path="workspace" element={<WorkspacePage />} />
                <Route path="ontology" element={<OntologyPage />} />
                <Route path="ontology_prop" element={<OntologyPropPage />} />
                <Route path="triple" element={<TriplePage />} />
                <Route path="extract_task" element={<ExtractTaskPage />} />
                <Route path="extract_task_result" element={<ExtractTaskResultPage />} />
                <Route path="knowledge_infer" element={<KnowledgeInferPage />} />

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
