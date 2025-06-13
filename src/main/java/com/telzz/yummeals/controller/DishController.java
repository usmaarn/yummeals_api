package com.telzz.yummeals.controller;

import java.util.List;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.telzz.yummeals.entity.Dish;
import com.telzz.yummeals.service.DishService;
import com.telzz.yummeals.service.impl.DishServiceImpl;
import org.springframework.web.bind.annotation.PutMapping;

@RestController
@RequestMapping("v1/dishes")
public class DishController {
    
    private DishService dishService;

    public DishController(DishServiceImpl dishService){
        this.dishService = dishService;
    }

    @GetMapping
    public List<Dish> findAll() {
        return dishService.findAll();
    }

    @PostMapping
    public ResponseEntity<String> addDish(@RequestBody Dish dish) {
        dishService.createDish(dish);
        return new ResponseEntity<String>("Dish added successfully!", HttpStatus.CREATED);
    }  

    @GetMapping("/{id}")
    public ResponseEntity<Dish> getDishById(@PathVariable("id") Long dishId) {
        Dish dish = dishService.getDishById(dishId);
        if(dish != null){
            return ResponseEntity.ok(dish);
        }
        return new ResponseEntity<>(HttpStatus.NOT_FOUND);
    }

    @PutMapping("/{id}")
    public ResponseEntity<String> updateDish(@PathVariable Long id, @RequestBody Dish dish) {
        this.dishService.updateDish(id, dish);
        return new ResponseEntity<String>("Dish updated successfully!", HttpStatus.OK);
    }

    @DeleteMapping("/{id}")
    public void deleteDishById(@PathVariable("id") Long dishId){
        dishService.deleteDishById(dishId);
    }
}