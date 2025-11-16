# Let's Go! React with Go Complete Fullstack App - TypeScript, React Query, MongoDB, ChakraUI

![Demo App](https://i.ibb.co/JvRTWmW/Group-93.png)

[Video Tutorial on Youtube](https://youtu.be/zw8z_o_kDqc)

Some Features:

-   ⚙️ Tech Stack: Go, React, TypeScript, MongoDB, TanStack Query, ChakraUI
-   ✅ Create, Read, Update, and Delete (CRUD) functionality for todos
-   🌓 Light and Dark mode for user interface
-   📱 Responsive design for various screen sizes
-   🌐 Deployment
-   🔄 Real-time data fetching, caching, and updates with TanStack Query
-   🎨 Stylish UI components with ChakraUI
-   ⏳ And much more!

### .env file

```shell
MONGO_URI=<your_mongo_uri>
PORT=5000
ENV=development
```

### Compile and run

```shell
go run main.go
```

# build stage
go build -o app ./

# start stage
./app

export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:5000/api" : "/api";
