package com.telzz.yummeals.service.impl;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import org.springframework.stereotype.Service;

import com.telzz.yummeals.entity.Dish;
import com.telzz.yummeals.service.DishService;

@Service
public class DishServiceImpl implements DishService {
    private List<Dish> dishes = new ArrayList<>();
    private Long nextId = 1L;

    @Override
    public List<Dish> findAll(){
        return this.dishes;
    }

    @Override
    public void createDish(Dish dish){
        dish.setId(nextId);
        this.dishes.add(dish);
        this.nextId++;
    }

    @Override
    public Dish getDishById(Long id){
        for (Dish dish : dishes){
            if(dish.getId().equals(id)){
                return dish;
            }
        }
        return null;
    }

    @Override
    public void updateDish(Long id, Dish dto){
        for(Dish dish : dishes){
            if(dish.getId().equals(id)){
                dish.setName(dto.getName());
                dish.setDescription(dto.getDescription());
                dish.setCategory(dto.getCategory());
                dish.setImage(dto.getImage());
            }
        }
    }

    @Override
    public void deleteDishById(Long id){
        for (int i = 0; i < dishes.size(); i ++){
            Dish dish = dishes.get(i);
            if(dish.getId().equals(id)){
                dishes.remove(i);
            }
        }
    }
}