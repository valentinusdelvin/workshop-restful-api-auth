# Case Study: Restaurant Menu Discovery API

## Tech Stack

- **Language:** Go 1.25.7
- **Web Framework:** Gin v1.11.0
- **Database:** PostgreSQL
- **ORM:** GORM v1.31.1
- **Database Driver:** gorm.io/driver/postgres v1.6.0
- **UUID:** github.com/google/uuid v1.6.0
- **Authentication:** github.com/golang-jwt/jwt
- **Password Hashing:** github.com/golang/crypto/bcrypt
- **Validation:** github.com/go-playground/validator/v10

## Overview

**FoodHub** is a mobile and web application that helps users discover and view restaurant menus in their city. Users can browse restaurants by location, see detailed menu items with prices, and check availability. Restaurant managers can add, update, and manage their menu items through the API.

This case study demonstrates a RESTful API for managing restaurant data and menu information. The system features two interconnected entities: **Restaurants** and **Menu Items**, allowing users to explore dining options without needing authentication for basic browsing. It showcases all CRUD operations in a real-world scenario.

## Data Models

### Restaurant Entity (GORM)
```go
type Restaurant struct {
    Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name      string    `gorm:"type:varchar(100);not null"`
    Location  string    `gorm:"type:varchar(50);not null"`
    CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
    Items     []Item    `gorm:"foreignKey:RestaurantId"`
}
```

### Item Entity (GORM)
```go
type Item struct {
    Id           uuid.UUID `gorm:"type:uuid;primaryKey"`
    RestaurantId uuid.UUID `gorm:"type:uuid;not null;constraint:OnDelete:CASCADE"`
    Name         string    `gorm:"type:varchar(100);not null"`
    Price        float64   `gorm:"type:decimal(10,2);not null"`
    Available    bool      `gorm:"type:boolean;not null;default:false"`
    CreatedAt    time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
    UpdatedAt    time.Time `gorm:"type:timestamp;not null;autoUpdateTime"`
}
```

## API Endpoints

### Restaurants

#### Create Restaurant (POST)
**Endpoint:** `POST /api/v1/restaurants`

**Description:** Create a new restaurant with a name and location, only user with role admin can use this endpoint.

**Request:**
```json
{
  "name": "Pasta House",
  "location": "Jakarta"
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Pasta House",
  "location": "Jakarta",
  "created_at": "2024-01-15T10:30:00Z"
}
```

---

#### Get All Restaurants (GET)
**Endpoint:** `GET /api/v1/restaurants?page=1&limit=10`

**Query Parameters:**
- `page` - Page number (default: 1)
- `limit` - Items per page (default: 10, max: 50)

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Pasta House",
      "location": "Jakarta",
      "created_at": "2024-01-15T10:30:00Z"
    },
    {
      "id": "550e8400-e29b-41d4-a716-446655440001",
      "name": "Sushi Place",
      "location": "Jakarta",
      "created_at": "2024-01-15T11:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10
  }
}
```

**Notes:**
- Results are ordered by `created_at DESC` (newest first)
- Omitting `page` and `limit` will use defaults
- Invalid page/limit values will return 400 Bad Request
- Only Authenticated users can access this endpoint resource

---

#### Update Restaurant (PATCH)
**Endpoint:** `PATCH /api/v1/restaurants/{id}`

**Request:**
```json
{
  "name": "Pasta House Premium",
  "location": "Jakarta"
}
```

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Pasta House Premium",
  "location": "Jakarta",
  "created_at": "2024-01-15T10:30:00Z"
}
```

---

#### Delete Restaurant (DELETE)
**Endpoint:** `DELETE /api/v1/restaurants/{id}`

**Response (204 No Content)**

**Notes:**
- Deleting a restaurant will cascade delete all associated items

---

### Items

#### Create Item (POST)
**Endpoint:** `POST /api/v1/restaurants/{restaurant_id}/items`

**Request:**
```json
{
  "name": "Carbonara",
  "price": 85000
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "restaurant_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Carbonara",
  "price": 85000,
  "available": false,
  "created_at": "2024-01-15T10:45:00Z",
  "updated_at": "2024-01-15T10:45:00Z"
}
```

---

#### Get All Items for Restaurant (GET)
**Endpoint:** `GET /api/v1/restaurants/{restaurant_id}/items?page=1&limit=10`

**Query Parameters:**
- `page` - Page number (default: 1)
- `limit` - Items per page (default: 10, max: 50)

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440002",
      "restaurant_id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Carbonara",
      "price": 85000,
      "available": false,
      "created_at": "2024-01-15T10:45:00Z",
      "updated_at": "2024-01-15T10:45:00Z"
    },
    {
      "id": "550e8400-e29b-41d4-a716-446655440003",
      "restaurant_id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Bolognese",
      "price": 75000,
      "available": false,
      "created_at": "2024-01-15T10:50:00Z",
      "updated_at": "2024-01-15T10:50:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10
  }
}
```

**Notes:**
- Results are ordered by `created_at DESC` (newest first)
- Returns 404 if restaurant does not exist

---

#### Update Item (PATCH)
**Endpoint:** `PATCH /api/v1/restaurants/{restaurant_id}/items/{item_id}`

**Request (Mark unavailable):**
```json
{
  "available": true
}
```

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "restaurant_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Carbonara",
  "price": 85000,
  "available": true,
  "created_at": "2024-01-15T10:45:00Z",
  "updated_at": "2024-01-15T11:30:00Z"
}
```

**Notes:**
- Only provided fields are updated
- Returns 404 if item does not exist

---

#### Delete Item (DELETE)
**Endpoint:** `DELETE /api/v1/restaurants/{restaurant_id}/items/{item_id}`

**Response (204 No Content)**

**Notes:**
- Returns 404 if item does not exist

---

## Database Schema

```sql
CREATE TABLE restaurants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    location VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    restaurant_id UUID NOT NULL REFERENCES restaurants(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    available BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```
