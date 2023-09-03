import { Routes, Route} from "react-router-dom";
// import { useState, useEffect } from "react";
// import FileView from "./components/FileView";
import FileView from "./components/FileView";
import GridView from "./components/GridView";

// const User = () => {
//   return (
//     <>
//       <h2>User</h2>

//       <nav>
//         <Link to="/user/profile">Profile</Link>
//         <Link to="/user/account">Account</Link>
//       </nav>
//     </>
//   );
// };

// const Users = ({ users }) => {
//   return (
//     <>
//       <h2>Users</h2>

//       <ul>
//         {users.map((user) => (
//           <li key={user.id}>
//             <Link to={user.id}>{user.fullName}</Link>
//           </li>
//         ))}
//       </ul>

//       <Outlet />
//     </>
//   );
// };

// const UsersPage = () => {
//   return (
//     <>
//       <h1>Users page</h1>

//       <Outlet />
//     </>
//   );
// };

// const File = () => {
//   return (
//     <>
//       <h2>User</h2>

//       <nav>
//         <Link to="/user/profile">Profile</Link>
//         <Link to="/user/account">Account</Link>
//       </nav>
//     </>
//   );
// };

// const Files = () => {
//   return (
//     <>
//       <h2>Files</h2>

//       <Outlet />
//     </>
//   );
// }

// const FilesPage = () => {
//   return (
//     <>
//       <h1>Files page</h1>

//       <Outlet />
//     </>
//   );
// }

const App = () => {
  // const users = [
  //   { id: "1", fullName: "Robin Wieruch" },
  //   { id: "2", fullName: "Sarah Finnley" }
  // ];

  

  return (
    <>
      <h1>React Router</h1>
    <GridView/>
      {/* <ul>
        <li>
          <Link to="/home">Home</Link>
        </li>
        <li>
          <Link to="/files"><GridView files={items}/></Link>
        </li>
      </ul>
  */}
      <Routes>
        <Route path="/" element={"Home"} />
        <Route path="/files" element={<GridView />}>
          <Route index element={<GridView />} />
          <Route path=":fileId" element={<FileView />} />
        </Route>
      </Routes> 
    </>
  );
};

export default App;



// import React from 'react';
// import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
// import GridView from './components/GridView';
// // import Home from './Home'; // Replace with your Home component
// // import About from './About'; // Replace with your About component
// // import Contact from './Contact'; // Replace with your Contact component
// // import NotFound from './NotFound'; // Replace with your Not Found (404) component

// function App() {
//   return (
//     <Router>
//       <div>
//         {/* Navigation */}
//         <nav>
//           <ul>
//             <li>
//               <Link to="/">Home</Link>
//             </li>
//             {/* <li>
//               <Link to="/about">About</Link>
//             </li>
//             <li>
//               <Link to="/contact">Contact</Link>
//             </li> */}
//           </ul>
//         </nav>

//         {/* Route Configuration */}
//         <Routes>
//           <Route path="/" element={<GridView />} />
//           {/* <Route path="/about" element={<About />} />
//           <Route path="/contact" element={<Contact />} />
//           <Route path="*" element={<NotFound />} /> Handle 404 */}
//         </Routes>
//       </div>
//     </Router>
//   );
// }

// export default App;



// import React, { Component } from 'react';
// import './App.css';
// // import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
// import GridView from './components/GridView';
// import FileView from './components/FileView';

// class App extends Component {
//   render() {
//       // return (
//       //   <div className='App'>
//       //   {/* <Router>
//       //     <Switch> */}
//       //       {/* <Route path="/" exact component={GridView} /> */}
//       //       {/* <Route path="/files/:fileId" component={FileView} />
//       //     </Switch>
//       //   </Router> */}
//       //   </div>
//       // );
    
//       return (
//       <div className="App">
        
//         <GridView/>
//         <FileView/>
//       </div>
//     );
//   }
// }

// export default App;
