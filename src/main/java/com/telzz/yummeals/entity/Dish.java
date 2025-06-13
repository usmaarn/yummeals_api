
package com.telzz.yummeals.entity;

public class Dish {
    private Long id;
    private String name;
    private String description;
    private String category;
    private String image;

    public void setId(Long id){
        this.id = id;
    }

    public Long getId(){
        return id;
    }

    public void setName(String name){
        this.name = name;
    }

    public String getName(){
        return name;
    }

    public void setDescription(String description){
        this.description = description;
    }

    public String getDescription(){
        return description;
    }

    public void setCategory(String category){
        this.category = category;
    }

    public String getCategory(){
        return category;
    }

    public void setImage(String image){
        this.image = image;
    }

    public String getImage(){
        return image;
    }
}