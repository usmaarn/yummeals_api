
package com.telzz.yummeals.service;

import java.util.List;

import com.telzz.yummeals.entity.Dish;

public interface DishService {
    List<Dish> findAll();
    void createDish(Dish dish);
    Dish getDishById(Long id);
    void deleteDishById(Long id);
    void updateDish(Long id, Dish dish);
}