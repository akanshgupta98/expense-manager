import React, { useState, useEffect } from "react";
import { BrowserRouter as Router, Routes, Route, Link, Navigate, useNavigate } from "react-router-dom";
import axios from "axios";

const API_BASE = "http://localhost:8081/api/v1"; // Change if your backend runs elsewhere

// Simple Auth Context to share user info
const AuthContext = React.createContext();

function useAuth() {
  return React.useContext(AuthContext);
}

function AuthProvider({ children }) {
  const [user, setUser] = useState(null);

  // Save token in memory for simplicity
  const login = (token) => {
    localStorage.setItem("token", token);
    setUser({ token });
  };
  const logout = () => {
    localStorage.removeItem("token");
    setUser(null);
  };

  // On mount check if token exists
  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) setUser({ token });
  }, []);

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

// Protected Route wrapper
function PrivateRoute({ children }) {
  const { user } = useAuth();
  return user ? children : <Navigate to="/login" />;
}

function NavBar() {
  const { user, logout } = useAuth();
  return (
    <nav className="bg-gray-800 p-4 flex justify-between items-center text-white">
      <Link to="/" className="font-semibold text-xl">Expense Manager</Link>
      <div className="space-x-4">
        {!user && (
          <>
            <Link to="/login" className="hover:underline">Login</Link>
            <Link to="/register" className="hover:underline">Register</Link>
          </>
        )}
        {user && (
          <>
            <Link to="/users" className="hover:underline">Users</Link>
            <button onClick={logout} className="hover:underline">Logout</button>
          </>
        )}
      </div>
    </nav>
  );
}

function Register() {
  const [form, setForm] = useState({ email: "", name: "",password: "" });
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);
  const navigate = useNavigate();

  const onChange = (e) => setForm({ ...form, [e.target.name]: e.target.value });

  const onSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);
    try {
      await axios.post(`${API_BASE}/user/users`, form);
      setSuccess("Registration successful! You can login now.");
      setForm({ email: "", name:"",password: "" });
      setTimeout(() => navigate("/login"), 1500);
    } catch (err) {
      setError(err.response?.data?.message || "Registration failed");
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl mb-6 font-semibold">Register</h2>
      {error && <p className="mb-4 text-red-600">{error}</p>}
      {success && <p className="mb-4 text-green-600">{success}</p>}
      <form onSubmit={onSubmit} className="space-y-4">
        <input
          name="name"
          value={form.name}
          onChange={onChange}
          placeholder="Name"
          required
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <input
          name="email"
          value={form.email}
          onChange={onChange}
          placeholder="Email"
          required
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <input
          name="password"
          type="password"
          value={form.password}
          onChange={onChange}
          placeholder="Password"
          required
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        
        
        <button
          type="submit"
          className="w-full bg-blue-600 hover:bg-blue-700 text-white py-2 rounded font-semibold"
        >
          Register
        </button>
      </form>
    </div>
  );
}

function Login() {
  const [form, setForm] = useState({ email: "", password: "" });
  const [error, setError] = useState(null);
  const { login } = useAuth();
  const navigate = useNavigate();

  const onChange = (e) => setForm({ ...form, [e.target.name]: e.target.value });

  const onSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    try {
      const res = await axios.post(`${API_BASE}/auth/login`, form);
      console.log(res,"RES");
      console.log(res.data.response.data.JWTToken,"TOKEN IS");
      login(res.data.response.data.JWTToken); // expecting {token: "jwt_token_here"}
      navigate("/users");
    } catch (err) {
      console.log("ERR IS",err)
      setError(err.response?.data?.message || "Login failed");
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl mb-6 font-semibold">Login</h2>
      {error && <p className="mb-4 text-red-600">{error}</p>}
      <form onSubmit={onSubmit} className="space-y-4">
        <input
          name="email"
          value={form.email}
          onChange={onChange}
          placeholder="Email"
          required
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <input
          name="password"
          type="password"
          value={form.password}
          onChange={onChange}
          placeholder="Password"
          required
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <button
          type="submit"
          className="w-full bg-blue-600 hover:bg-blue-700 text-white py-2 rounded font-semibold"
        >
          Login
        </button>
      </form>
    </div>
  );
}

function Users() {
  const [users, setUsers] = useState([]);
  const [error, setError] = useState(null);
  const { user } = useAuth();

  useEffect(() => {
    async function fetchUsers() {
      setError(null);
      try {
        const res = await axios.get(`${API_BASE}/user/users`, {
          headers: { Authorization: `Bearer ${user.token}` },
        });
        setUsers(res.data.Users);
      } catch (err) {
        setError(err.response?.data?.message || "Failed to load users");
      }
    }
    fetchUsers();
  }, [user]);

  return (
    <div className="max-w-3xl mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl mb-6 font-semibold">Users</h2>
      {error && <p className="mb-4 text-red-600">{error}</p>}
      {!error && users.length === 0 && <p>No users found.</p>}
      <ul className="divide-y divide-gray-200">
        {users.map((u) => (
          <li key={u.id} className="py-3 flex justify-between">
            <span>{u.username}</span>
            <span className="text-gray-500 text-sm">{u.email || "-"}</span>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default function App() {
  return (
    <AuthProvider>
      <Router>
        <div className="min-h-screen bg-gray-100">
          <NavBar />
          <Routes>
            <Route path="/" element={<Navigate to="/login" />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route
              path="/users"
              element={
                <PrivateRoute>
                  <Users />
                </PrivateRoute>
              }
            />
            <Route path="*" element={<p className="p-6">404 Not Found</p>} />
          </Routes>
        </div>
      </Router>
    </AuthProvider>
  );
}
