package org.fi;

import java.io.Serializable;


//Serialization in Java allows us to convert an Object to stream that we can send over 
//the network or save it as file or store in DB for later usage.

public class Alert implements Serializable{
    private long id;
    private String message;

    public Alert() {
    }

    public Alert(long id, String message) {
        this.id = id;
        this.message = message;
    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    @Override
    public String toString() {
        return "Alert{" +
                "id=" + id +
                ", message='" + message + '\'' +
                '}';
    }
}
